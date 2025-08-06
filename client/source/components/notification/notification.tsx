import { Box, Text } from "ink";
import BigText from "ink-big-text";
import Gradient from "ink-gradient";
import React from "react";

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
