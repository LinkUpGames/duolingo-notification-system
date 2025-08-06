import { Box } from "ink";
import React from "react";
import Notify, { type Notification } from "../notification/notification.js";

/**
 * The main area
 */
const Main = () => {
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
