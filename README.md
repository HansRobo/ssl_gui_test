# ssl_gui_test

## Setup and Startup

### Frontend (Vue.js)

1. Navigate to the `frontend` directory:
   ```bash
   cd frontend
   ```

2. Install the dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run serve
   ```

### Backend (Golang)

1. Navigate to the `backend` directory:
   ```bash
   cd backend
   ```

2. Install the dependencies:
   ```bash
   go mod tidy
   ```

3. Start the backend server:
   ```bash
   go run main.go
   ```

### Usage

1. Open your browser and navigate to `http://localhost:8080`.
2. You should see a canvas where shapes will be drawn based on the data received from the backend.
3. The backend will forward UDP packets to the frontend via WebSocket, and the shapes will be drawn on the canvas accordingly.

## Developing with Devcontainer

1. Open the repository in Visual Studio Code.
2. Install the "Remote - Containers" extension if you haven't already.
3. Press `F1` and select `Remote-Containers: Open Folder in Container...`.
4. Select the root folder of the repository.
5. The development container will be built and started automatically.
6. You can now develop the project inside the container.
