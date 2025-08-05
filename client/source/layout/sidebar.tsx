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
      borderStyle="double"
    ></Box>
  );
};

export default Sidebar;
