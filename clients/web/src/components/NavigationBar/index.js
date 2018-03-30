import React from "react";
import Media from "react-media";
import { withTheme } from "styled-components";

import Box from "../Box";

const NavigationBar = ({ theme: { breakpoints } }) => {
    return (
        <Media query={`(min-width: ${breakpoints[1]})`}>
            {matches =>
                matches ? (
                    <Box bg="wireframe" py={3} px={3}>
                        <div>Test</div>
                    </Box>
                ) : null
            }
        </Media>
    );
};

export default withTheme(NavigationBar);
