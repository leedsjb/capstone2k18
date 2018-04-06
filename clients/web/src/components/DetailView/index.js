import React from "react";
import Media from "react-media";
import { Flex } from "grid-styled";
import { withTheme } from "styled-components";

import Box from "../Box";

const DetailView = ({ children, theme: { breakpoints } }) => {
    return (
        <Media
            query={`(min-width: ${breakpoints[1]})`}
            render={() => {
                return (
                    <Flex flex={1}>
                        <Box borderLeft="1px solid" width={1 / 1} height="100%">
                            {children}
                        </Box>
                    </Flex>
                );
            }}
        />
    );
};

export default withTheme(DetailView);
