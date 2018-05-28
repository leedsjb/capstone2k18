import React from "react";

import Box from "../Box";
import Span from "../Span";
import Text from "../Text";

const Requestor = ({ requestor }) => {
    return (
        <div>
            {requestor.name ? (
                <Box mt={4}>
                    <Span>{requestor.name}</Span>
                </Box>
            ) : null}
            {requestor.phone ? (
                <Box mt={4}>
                    <Span>
                        {requestor.phone.length === 10
                            ? `(${requestor.phone.substring(
                                  0,
                                  3
                              )}) ${requestor.phone.substring(
                                  3,
                                  6
                              )}-${requestor.phone.substring(6, 10)}`
                            : requestor.phone}
                    </Span>
                </Box>
            ) : null}
            {requestor.type ? (
                <Box mt={4}>
                    <Span>{requestor.type}</Span>
                </Box>
            ) : null}
            {requestor.address &&
            requestor.city &&
            requestor.state &&
            requestor.zip ? (
                <Box mt={4}>
                    <Text mt={1}>{requestor.address}</Text>
                    <Text>{`${requestor.city}, ${requestor.state} ${
                        requestor.zip
                    }`}</Text>
                </Box>
            ) : null}
        </div>
    );
};

export default Requestor;
