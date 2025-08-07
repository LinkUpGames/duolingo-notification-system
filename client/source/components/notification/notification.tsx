import { Box, Text } from "ink";
import BigText from "ink-big-text";
import Gradient from "ink-gradient";
import React from "react";
import Logger from "../../logs.js";

/**
 * The notification gathered from the backend
 */
export type Notification = {
  /**
   * The title of the notification
   */
  title: string;

  /**
   * Text about the notificaiton
   */
  description: string;

  /**
   * Notificaiton id
   */
  id: string;

  /**
   * The score for the notification
   */
  score: number;

  /**
   * The millisecond timestamp
   */
  timestamp: string;

  /**
   * The probability that this notificaiton was selected
   */
  probability: number;
};

type Props = {
  /**
   * The notification to render
   */
  notification: Notification;
};

/**
 * Notification to display
 */
const Notification = ({ notification }: Props) => {
  Logger.log({
    level: "info",
    message: `Notification: ${JSON.stringify(notification)}`,
  });

  return (
    <Box
      borderStyle="round"
      borderColor="cyan"
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
    >
      <Gradient name="instagram">
        <BigText text={notification.title} />
      </Gradient>

      <Text dimColor color="magentaBright">
        {notification.description}
      </Text>
    </Box>
  );
};

export default Notification;
