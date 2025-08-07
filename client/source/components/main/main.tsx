import { Box, Text } from "ink";
import React, { useContext, useEffect, useState } from "react";
import Notify, { type Notification } from "../notification/notification.js";
import { AppContext, useFetch } from "../../utils.js";

/**
 * The main area
 */
const Main = () => {
  const { server } = useContext(AppContext);
  const { fetch, loading, data } = useFetch<Notification>(server); // Fetch the notification

  // STATES
  const [notification, setNotification] = useState<Notification | null>(null); // The notification fetched from the backend

  // FUNCTIONS
  /**
   * Fetch a notification from the database based on the algorithm
   */
  const fetchNotification = async () => {
    try {
      const result = await fetch("send_notification?user_id=1233");

      setNotification(result);
    } catch {}
  };

  // EFFECTS
  /**
   * On an initial render, fetch the notification data first
   */
  useEffect(() => {
    fetchNotification();
  }, []);

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
      {notification ? (
        <Notify notification={notification} />
      ) : (
        <Text> No Notification</Text>
      )}
    </Box>
  );
};

export default Main;
