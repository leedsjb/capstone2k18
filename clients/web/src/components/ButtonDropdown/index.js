import React from "react";
import { Flex } from "grid-styled";

import Icon from "../Icon";
import Box from "../Box";
import Button from "../Button";

const ButtonDropdown = ({ children, ...props }) => {
    return (
        <Button py={2} px={3} bg="green" {...props}>
            <Flex alignItems="center">
                <Box mr={2}>{children}</Box>
                <Icon glyph="triangleDown" size={6} />
            </Flex>
        </Button>
    );
};

export default ButtonDropdown;
