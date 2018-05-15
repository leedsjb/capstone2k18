import React from "react";
import { Flex } from "grid-styled";

import Button from "../Button";
import Icon from "../Icon";
import Box from "../Box";

const ButtonDropdown = ({ children, ...props }) => {
    return (
        <Button py={1} px={2} {...props}>
            <Flex alignItems="center">
                <Box mr={2}>{children}</Box>
                <Icon glyph="triangleDown" size={6} />
            </Flex>
        </Button>
    );
};

export default ButtonDropdown;
