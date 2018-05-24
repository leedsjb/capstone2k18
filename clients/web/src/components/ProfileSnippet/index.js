import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import Span from "../Span";

const ProfileSnippet = ({ label, value, ...props }) => {
    return (
        <Flex flexDirection="column" alignItems="center" {...props}>
            <Box>
                <Span fontWeight="bold" textAlign="center">
                    {label}
                </Span>
            </Box>
            <Box mt={1}>
                <Span textAlign="center">{value}</Span>
            </Box>
        </Flex>
    );
};

export default ProfileSnippet;
