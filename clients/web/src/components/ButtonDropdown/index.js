import React from "react";
import { Flex } from "grid-styled";

import Icon from "../Icon";
import Box from "../Box";
import ButtonPrimary from "../ButtonPrimary";

const ButtonDropdown = ({ children, ...props }) => {
    return (
        <ButtonPrimary py={2} px={3} {...props}>
            <Flex alignItems="center">
                <Box mr={2}>{children}</Box>
                <Icon glyph="triangleDown" size={6} />
            </Flex>
        </ButtonPrimary>
    );
};

export default ButtonDropdown;
