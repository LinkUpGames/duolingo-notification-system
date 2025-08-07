import React from "react";
import { AppContext, useSetup } from "./utils.js";
import Window from "./layout/window.js";
import Main from "./components/main/main.js";
import Logger from "./logs.js";
import { useApp, useInput } from "ink";

type Props = {
  name: string | undefined;
};

export default function App({ name = "Stranger" }: Props) {
  const { server } = useSetup();

  Logger.info("Name", { message: name });

  const { exit } = useApp();

  // Key inputs
  useInput((input, key) => {
    if (key.escape || input === "q") {
      exit();
    }
  });

  return (
    <AppContext.Provider
      value={{
        server: server,
      }}
    >
      {/* Global Window */}
      <Window width="100%" height="100%">
        {/* <Sidebar width="30%" /> */}

        <Main />
      </Window>
    </AppContext.Provider>
  );
}
