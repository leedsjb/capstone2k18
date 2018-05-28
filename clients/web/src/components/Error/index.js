import React from "react";
import { Flex } from "grid-styled";
import { withTheme } from "styled-components";
import { Link } from "react-router-dom";

import Text from "../Text";
import Span from "../Span";
import Heading from "../Heading";
import Box from "../Box";
import Illustration from "../Illustration";
import Button from "../Button";

import PushProvider from "../../containers/PushProvider";

const Error = ({ title, content, theme: { colors }, notFound }) => {
    if (notFound) {
        title = "TEST";
        content = "Placeholder text lalala";
    }

    return (
        <Flex flexDirection="column" alignItems="center">
            <Illustration glyph="airport" size={160} color={colors.black1} />
            <Heading is="h3" fontSize={4} textAlign="center" mt={2}>
                {title}
            </Heading>
            <Heading
                is="h4"
                fontSize={2}
                fontWeight="normal"
                textAlign="center"
                mt={1}
            >
                {content}
            </Heading>
            <Box mt={6}>
                {notFound ? (
                    <Button onClick={() => window.location.reload()}>
                        Refresh page
                    </Button>
                ) : (
                    <Button is={Link} to="/aircraft">
                        Take me home
                    </Button>
                )}
            </Box>
        </Flex>
    );
};

export default withTheme(Error);