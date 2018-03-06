import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import Heading from "../Heading";

const TitleBar = () => {
    return (
        <Box bg="wireframe" py={3} px={3}>
            <Flex justifyContent="space-between" align="center">
                <Box size={32} bg="black" />
                <Heading is="h3" fontSize={2}>Title</Heading>
                <Box size={32} bg="black" />
            </Flex>
        </Box>
    );
};

export default TitleBar;