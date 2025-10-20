/*
Basic DDD example in Go

This application intends to provide a simple demonstration of key architecture and design concepts of modern software in an easily consumable format
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
----------------------------------------------------------------

	DOMAIN LAYER

	- The core business logic and rules are encapsulated in the domain layer
----------------------------------------------------------------
*/

// our domain entity for PingPong is represented as a struct with a message
type PingPongEntity struct {
	message string // can be "ping" or "pong"
}

// validate ensures the entity is in a valid state
func (e *PingPongEntity) validate() error {
	if e.message != "ping" && e.message != "pong" {
		return errors.New("ya gotta send a ping or a pong") // this error would be handled in the service layer
	}

	return nil
}

// GetMessage returns the message of the entity
func (e *PingPongEntity) GetMessage() string {
	return e.message
}

func (e *PingPongEntity) DetermineResult() (*PingPongEntity, error) {
	if e.message == "ping" {
		return &PingPongEntity{message: "PONG!"}, nil
	}
	if e.message == "pong" {
		return &PingPongEntity{message: "PING!"}, nil
	}
	err := errors.New("bummer")
	return nil, err
}

// End Domain -----------------------------------------------------------------

/*
----------------------------------------------------------------

	INTERFACE LAYER

----------------------------------------------------------------
*/
type IPingPongCommandLineInterface interface {
	HandleRequest(request *CommandLineInterfacePingPongRequest) (*CommandLineInterfacePingPongResponse, error)
}

type pingPongCommandLineInterface struct {
	svc *PingPongService
}

func NewCommandLineInterfacePingPongRequest(input string) *CommandLineInterfacePingPongRequest {
	sanitizedInput := strings.TrimSpace(input)
	return &CommandLineInterfacePingPongRequest{sanitizedInput}
}

type CommandLineInterfacePingPongRequest struct {
	message string
}

type CommandLineInterfacePingPongResponse struct {
	message string
}

func (r *CommandLineInterfacePingPongResponse) String() string {
	return r.message
}

func NewPingPongCommandLineInterface(svc *PingPongService) *pingPongCommandLineInterface {
	return &pingPongCommandLineInterface{svc: svc}
}

func (i *pingPongCommandLineInterface) HandleRequest(request *CommandLineInterfacePingPongRequest) (*CommandLineInterfacePingPongResponse, error) {
	if request == nil {
		return nil, errors.New("empty request")
	}

	service := NewPingPongService(NewPingPongStdOutRepo())
	cmd := MapPingPongRequestToCommand(request)

	result, err := service.PingPong(cmd)
	if err != nil {
		return nil, err
	}

	return MapPingPongCommandToResponse(result), nil
}

// End Interface -----------------------------------------------------------------

// ----------------------------------------
// INTERFACE -> SERVICE Translations
// ----------------------------------------

func MapPingPongRequestToCommand(request *CommandLineInterfacePingPongRequest) *PingPongCommand {
	return &PingPongCommand{message: request.message}
}

func MapPingPongCommandToResponse(command *PingPongCommandResult) *CommandLineInterfacePingPongResponse {
	return &CommandLineInterfacePingPongResponse{message: command.message}
}

/* End of Translation Layer */

/*
----------------------------------------------------------------

	SERVICE LAYER

----------------------------------------------------------------
*/
type IPingPongService interface {
	PingPong(cmd *PingPongCommand) (*PingPongCommandResult, error)
}

type PingPongService struct {
	repo IPingPongRepository
}

func NewPingPongService(repo IPingPongRepository) *PingPongService {
	return &PingPongService{
		repo: repo,
	}
}

type PingPongCommand struct {
	message string
}

type PingPongCommandResult struct {
	message string
}

func (s *PingPongService) PingPong(cmd *PingPongCommand) (*PingPongCommandResult, error) {
	pingPongReceivedEntity := &PingPongEntity{
		message: cmd.message,
	}

	if err := pingPongReceivedEntity.validate(); err != nil {
		return nil, err
	}

	pingPongResultingEntity, err := pingPongReceivedEntity.DetermineResult()
	if err != nil {
		return nil, err
	}

	if err := s.repo.SavePingPong(pingPongReceivedEntity, pingPongResultingEntity); err != nil {
		return nil, err
	}

	// Create the result to return
	result := &PingPongCommandResult{
		message: pingPongResultingEntity.GetMessage(),
	}

	return result, nil
}

// End Service -----------------------------------------------------------------

/*
----------------------------------------------------------------

	REPOSITORY LAYER

----------------------------------------------------------------
*/
type IPingPongRepository interface {
	SavePingPong(received *PingPongEntity, returning *PingPongEntity) error
}

type PingPongRepository struct {
	infrastructureConnection io.Writer
}

func NewPingPongStdOutRepo() *PingPongRepository {
	return &PingPongRepository{
		infrastructureConnection: os.Stdout,
	}
}

func (r *PingPongRepository) SavePingPong(received *PingPongEntity, returning *PingPongEntity) error {
	var err error
	_, err = fmt.Fprintf(r.infrastructureConnection, "saving to repo: %s\n", returning.message)
	if err != nil {
		return err
	}
	return nil
}

// End Repository -----------------------------------------------------------------

// Application Server struct defines the components of our application server
// it consists of an interface layer and an input/output mechanism
type ApplicationServer struct {
	interfaceLayer IPingPongCommandLineInterface
	inputOutput    *bufio.Reader // pretend this was a database connection?
}

// NewAppServer initializes the application server with its dependencies
func NewAppServer() *ApplicationServer {
	// Initialize the repository
	repo := NewPingPongStdOutRepo()

	// Initialize the service
	// the repository dependency is injected into the service
	svc := NewPingPongService(repo)

	// initialize the interface
	// the service dependency is injected into the interface
	interf := NewPingPongCommandLineInterface(svc)

	// bufio reader serves our interfaces
	server := bufio.NewReader(os.Stdin)

	return &ApplicationServer{
		interfaceLayer: interf,
		inputOutput:    server,
	}
}

// the ListenForRequest method is a continuous loop
func (a *ApplicationServer) ListenForRequest() {
	for {
		fmt.Print("> ") // prompt for input

		// read input from stdin until newline
		input, err := a.inputOutput.ReadString('\n')
		if err != nil { // if we hit an error, we break the loop
			break
		}

		req := NewCommandLineInterfacePingPongRequest(input)

		response, err := a.interfaceLayer.HandleRequest(req)
		if err != nil {
			fmt.Println(err)
			break // on error, we break the loop
		}

		fmt.Println("received response:", response.String())
		fmt.Println() // print a newline for better readability
	}
}

// main loop
func main() {
	fmt.Println("will you ping? or pong?") // prompt

	app := NewAppServer()
	app.ListenForRequest()
}
