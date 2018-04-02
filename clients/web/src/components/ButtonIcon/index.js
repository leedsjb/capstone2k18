import React from "react";
import { Flex } from "grid-styled";

import Button from "../Button";
import Icon from "../Icon";
import Box from "../Box";

const ButtonIcon = ({ glyph, children }) => {
    return (
        <Button>
            <Flex alignItems="center">
                <Icon glyph={glyph} size={16} />
                <Box ml={2}>{children}</Box>
            </Flex>
        </Button>
    );
};

export default ButtonIcon;
