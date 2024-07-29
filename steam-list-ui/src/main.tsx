import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { App } from "./Pages/App";
import { GameList } from "./Pages/GameList";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <App>
      <GameList />
    </App>
  </React.StrictMode>,
);
