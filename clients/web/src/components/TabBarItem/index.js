import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";

const TabBarItem = ({ title, onClick }) => {
    return (
        <Box bg="wireframe" flex={1} onClick={onClick}>
            <Flex justifyContent="center">
                <Flex flexDirection="column">
                    <Box p={3} bg="black" />
                    {title}
                </Flex>
            </Flex>
        </Box>
    )
};

export default TabBarItem;