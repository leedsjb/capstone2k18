import React from "react";
import { Flex } from "grid-styled";
import { withTheme } from "styled-components";
import { Link } from "react-router-dom";

import Heading from "../Heading";
import Box from "../Box";
import Illustration from "../Illustration";
import ButtonPrimary from "../ButtonPrimary";

const Error = ({ title, content, theme: { colors }, notFound }) => {
    if (notFound) {
        title = "Page Not Found";
        content = "The page you are looking for cannot be found";
    }

    return (
        <Flex flexDirection="column" alignItems="center">
            <Illustration glyph="airport" size={160} color={colors.black1} />
            <Heading is="h3" fontSize={4} textAlign="center" mt={4}>
                {title}
            </Heading>
            <Heading
                is="h4"
                fontSize={2}
                fontWeight="normal"
                textAlign="center"
                mt={2}
            >
                {content}
            </Heading>
            <Box mt={8}>
                {notFound ? (
                    <Link to="/aircraft">
                        <ButtonPrimary>Take me home</ButtonPrimary>
                    </Link>
                ) : (
                    <ButtonPrimary onClick={() => window.location.reload()}>
                        Refresh page
                    </ButtonPrimary>
                )}
            </Box>
        </Flex>
    );
};

export default withTheme(Error);
