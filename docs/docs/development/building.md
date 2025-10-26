# Building

The backend (core) application is packaged into a single binary, which references the built frontend files

The frontend application is packaged into relevant build files utilizing Vite, output to the `/dist/frontend` directory.

The backend application serves these files according to the logic defined within `/api/frontend/frontend.go`

After the backend build process completes, a Make `copier` is ran. This command copies the contents of `./dist/frontend` to all dist directories, enabling each compiled binary to reference the frontend files at *their* `./frontend`

![build](../_img/build.drawio.png)
