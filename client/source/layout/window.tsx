import { Box } from "ink";
import React from "react";
import { JSX } from "react";

type Props = {
  /**
   * Inner component
   */
  children: JSX.Element[] | JSX.Element;

  /**
   * The width of the window
   */
  width: number | string | undefined;

  /**
   * Height of the window
   */
  height: number | string | undefined;
};

/**
 * A window that fills the whole screen
 */
const Window = ({ children, width, height }: Props) => {
  return (
    <Box width={width} height={height} padding={1} borderStyle="single">
      {children}
    </Box>
  );
};

export default Window;
