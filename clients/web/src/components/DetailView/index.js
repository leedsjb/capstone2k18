import React from "react";
import Media from "react-media";
import { withTheme } from "styled-components";

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
