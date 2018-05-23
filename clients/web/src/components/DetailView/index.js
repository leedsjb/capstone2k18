import React from "react";
import Media from "react-media";
import { Flex } from "grid-styled";
import { withTheme } from "styled-components";

import Box from "../Box";
import ScrollView from "../ScrollView";

const DetailView = ({ children, theme: { breakpoints, colors } }) => {
    return (
        <Media
            query={`(min-width: ${breakpoints[1]})`}
            render={() => {
                return (
                    <ScrollView borderLeft={`1px solid ${colors.border}`}>
                        {children}
                    </ScrollView>
                );
            }}
        />
    );
};

export default withTheme(DetailView);
