import axios from "axios";
import { createContext, useMemo, useState } from "react";
import Logger from "./logs.js";

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

    return `http://server:${port}`;
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
export const useFetch = <T = unknown,>(host: string) => {
  // STATES
  const [loading, setLoading] = useState<boolean>(false); // controll whether the user has requested a fetch or not
  const [data, setData] = useState<T | null>(null); // The data that we fetched

  // FUNCTIONS
  /**
   * Fetch information from the backend based on the route we want
   */
  const fetch = async (route: string): Promise<T | null> => {
    let data: T | null = null;

    const url = `${host}/${route}`;

    try {
      setLoading(true);
      const response = await axios.get<string>(url);
      setLoading(false);
      setData(response.data as T);

      data = response.data as T;

      Logger.info("Data", { message: response.data });
    } catch (error) {
      Logger.log({ level: "error", message: error as string });
    }

    return data;
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
