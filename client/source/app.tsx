import React from "react";
import { AppContext, useSetup } from "./utils.js";
import Window from "./layout/window.js";
import Sidebar from "./layout/sidebar.js";
import Main from "./components/main/main.js";

type Props = {
  name: string | undefined;
};

export default function App({ name = "Stranger" }: Props) {
  const { server } = useSetup();

  console.log("Name: ", name);

  return (
    <AppContext.Provider
      value={{
        server: server,
      }}
    >
      {/* Global Window */}
      <Window width="100%" height="100%">
        <Sidebar width="30%" />

        <Main />
      </Window>
    </AppContext.Provider>
  );
}
