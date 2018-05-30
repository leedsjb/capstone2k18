import React from "react";
import { Flex } from "grid-styled";
import { withTheme } from "styled-components";

import Heading from "../Heading";
import Illustration from "../Illustration";

const EmptyState = ({ showContent, page, theme: { colors } }) => {
    let glyph;
    let size;
    let title;
    let content;

    switch (page) {
        case "aircraft":
            glyph = "balloon";
            size = 160;
            title = "No aircraft found";
            content = "Try searching for something else.";
            break;
        case "people":
            glyph = "pilot";
            size = 240;
            title = "No people found";
            content = "Try searching for something else.";
            break;
        case "groups":
            glyph = "pilot";
            size = 240;
            title = "No groups found";
            content = "Try searching for something else.";
            break;
        default:
            return;
    }

    return (
        <Flex flexDirection="column" alignItems="center">
            <Illustration glyph={glyph} size={size} color={colors.black1} />
            <Heading is="h3" fontSize={4} textAlign="center" mt={2}>
                {title}
            </Heading>
            <Heading
                is="h4"
                fontSize={2}
                fontWeight="normal"
                textAlign="center"
                mt={2}
            >
                {showContent ? content : null}
            </Heading>
        </Flex>
    );
};

export default withTheme(EmptyState);
