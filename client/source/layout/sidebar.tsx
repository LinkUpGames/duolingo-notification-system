import { Box } from "ink";
import React from "react";

type Props = {
  /**
   * The width of the container
   */
  width: number | string;
};

const Sidebar = ({ width }: Props) => {
  return (
    <Box
      height="100%"
      width={width}
      borderColor="green"
      borderStyle="round"
    ></Box>
  );
};

export default Sidebar;
