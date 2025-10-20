import Swal from "sweetalert2";

// showCreatePingPongResponse displays a notification with the created ping pong message
export function useShowCreatePingPongResponse(
  createdPingPongName: string | undefined,
) {
  let backgroundColor = "#1084d0"; // Default color

  if (createdPingPongName === "Ping!") {
    backgroundColor = "#e11d48"; // Red for ping
  } else if (createdPingPongName === "Pong!") {
    backgroundColor = "#0ea5e9"; // Blue for pong
  }

  // Display the notification using SweetAlert2
  Swal.fire({
    position: "bottom-end",
    html: `${createdPingPongName}`,
    timer: 1000,
    timerProgressBar: true,
    backdrop: false,
    showConfirmButton: false,
    toast: true,
    background: backgroundColor,
    animation: false,
    color: "#fff",
  });
}
