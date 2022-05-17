import { StrictMode } from "react";
import ReactDom from "react-dom";
import { CookiesProvider } from "react-cookie";

import App from "./App";

ReactDom.render(
  <StrictMode>
    <CookiesProvider>
      <App />
    </CookiesProvider>
  </StrictMode>,
  document.getElementById("root")
);