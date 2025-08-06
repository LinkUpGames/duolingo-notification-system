import { Box } from "ink";
import React, { useContext, useEffect } from "react";
import Notify, { type Notification } from "../notification/notification.js";
import { AppContext, useFetch } from "../../utils.js";

/**
 * The main area
 */
const Main = () => {
  const { server } = useContext(AppContext);
  const { fetch, loading, data } = useFetch(server);

  /**
   * On an initial render, fetch the notification data first
   */
  useEffect(() => {}, []);

  const notification: Notification = {
    title: "This is a title",
    description: "This is a description",
  };

  return (
    <Box
      flexDirection="column"
      overflow="visible"
      width="100%"
      height="99%"
      padding={1}
      borderStyle="round"
      borderColor="red"
    >
      <Notify notification={notification} />
    </Box>
  );
};

export default Main;
