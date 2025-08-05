import axios from "axios";
import { Box, Text } from "ink";
import React, {
  createContext,
  JSX,
  useContext,
  useMemo,
  useState,
} from "react";

type AppContextProps = {
  /**
   * The url for the server
   */
  server: string;
};

export const AppContext = createContext<AppContextProps>({
  server: "8080",
});

/**
 * Setup the client frontend based on what we are doing
 */
export const useSetup = () => {
  const { server: host } = useContext(AppContext);

  /**
   * The url for the server
   */
  const server = useMemo(() => {
    const port = process.env["SERVER_PORT"];

    return `http://localhost:${port}`;
  }, []);

  return {
    /**
     * The server to return
     */
    server,
  };
};

/**
 * Expose api to fetch and post notification stuff
 */
export const useFetch = (host: string) => {
  // STATES
  const [fetched, setFetched] = useState<boolean>(false); // controle whether the user has requested a fetch or not

  // FUNCTIONS
  /**
   * Fetch information from the backend based on the route we want
   */
  const fetch = async (route: string) => {
    const url = `${host}/${route}`;
    try {
      const response = await axios.get<string>(url);

      console.log("Data: ", response.data);
    } catch (error) {
      console.error("Error fetching: ", error);
    }
  };
};
