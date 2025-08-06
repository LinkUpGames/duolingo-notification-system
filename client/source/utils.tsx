import axios from "axios";
import { createContext, useMemo, useState } from "react";

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
  const [loading, setLoading] = useState<boolean>(false); // controll whether the user has requested a fetch or not
  const [data, setData] = useState<unknown | null>(null); // The data that we fetched

  // FUNCTIONS
  /**
   * Fetch information from the backend based on the route we want
   */
  const fetch = async (route: string) => {
    const url = `${host}/${route}`;
    try {
      setLoading(true);
      const response = await axios.get<string>(url);
      setLoading(false);
      setData(response.data);

      console.log("Data: ", response.data);
    } catch (error) {
      console.error("Error fetching: ", error);
    }
  };

  return {
    /**
     * Fetch information rom the bakend bvased on the route we want
     */
    fetch: fetch,

    /**
     * Whether the data is waiting to loading or not
     */
    loading: loading,

    /**
     * The data that was fetched
     */
    data: data,
  };
};
