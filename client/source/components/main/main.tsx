import { Box, useInput } from "ink";
import { DateTime } from "luxon";
import React, { useContext, useEffect, useState } from "react";
import Notify, { type Notification } from "../notification/notification.js";
import { AppContext, useFetch } from "../../utils.js";
import Logger from "../../logs.js";
import BigText from "ink-big-text";

/**
 * The main area
 */
const Main = () => {
  // Context
  const { server } = useContext(AppContext);
  const { fetch: getNotification, loading: getNotificationLoad } =
    useFetch<Notification>(server); // Fetch the notification that was selected to be shown
  const { fetch: postNotification } = useFetch<string>(server); // Post whether the notification worked or not

  // STATES
  const [notification, setNotification] = useState<Notification | null>(null); // The notification fetched from the backend

  // FUNCTIONS
  /**
   * Fetch a notification from the database based on the algorithm
   */
  const fetchNotification = async () => {
    try {
      const result = await getNotification("send_notification?user_id=1233");

      setNotification(result);
    } catch (error) {
      Logger.log({
        level: "error",
        message: `Error with fetching notification: ${error}`,
      });
    }
  };

  /**
   * Accept the notification that was sent
   * @param accepted Whether the notification was accepted or declined
   */
  const acceptNotification = async (accepted: boolean = false) => {
    try {
      await postNotification(
        `accept_notification?decision_id=${
          notification?.decision_id ?? ""
        }&selected=${accepted}&timestamp=${DateTime.now().toMillis()}`
      );

      fetchNotification();
    } catch (error) {
      Logger.log({
        level: "error",
        message: `Error with accepting the notification: ${error}`,
      });
    }
  };

  // EFFECTS
  /**
   * On an initial render, fetch the notification data first
   */
  useEffect(() => {
    fetchNotification();
  }, []);

  useInput(async (input, keys) => {
    if (keys.return) {
      // Accept the notification and update the value
      if (notification) {
        await acceptNotification(true);
      }
    }

    if (input === "n") {
      if (notification) {
        await acceptNotification(false);
      }
    }
  });

  return (
    <Box
      flexDirection="column"
      overflow="visible"
      width="100%"
      height="99%"
      padding={1}
      borderStyle="round"
      borderColor="red"
      flexGrow={1}
    >
      {!getNotificationLoad && notification ? (
        <Notify notification={notification} />
      ) : (
        <BigText text="No Notification" />
      )}
    </Box>
  );
};

export default Main;
