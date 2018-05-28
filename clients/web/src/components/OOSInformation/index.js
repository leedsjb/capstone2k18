import React from "react";

import Box from "../Box";
import Span from "../Span";

const OOSInformation = ({ OOS }) => {
    return (
        <div>
            {OOS.reason ? (
                <Box mt={4}>
                    <Span fontWeight="bold">OOS Reason: </Span>
                    <Span>{OOS.reason}</Span>
                </Box>
            ) : null}
            {OOS.remaining ? (
                <Box mt={4}>
                    <Span fontWeight="bold">Time Remaining: </Span>
                    <Span>{OOS.remaining}</Span>
                </Box>
            ) : null}
            {OOS.duration ? (
                <Box mt={4}>
                    <Span fontWeight="bold">OOS Duration: </Span>
                    <Span>{OOS.duration}</Span>
                </Box>
            ) : null}
        </div>
    );
};

export default OOSInformation;
