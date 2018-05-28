import React from "react";

import Box from "../Box";
import Span from "../Span";
import Text from "../Text";

const Receiver = ({ receiver }) => {
    return (
        <div>
            {receiver.name ? (
                <Box mt={4}>
                    <Span fontWeight="bold">Receiver Name: </Span>
                    <Span>{receiver.name}</Span>
                </Box>
            ) : null}
            {receiver.phone ? (
                <Box mt={4}>
                    <Span fontWeight="bold">Receiver Phone: </Span>
                    <Span>
                        {receiver.phone.length === 10
                            ? `(${receiver.phone.substring(
                                  0,
                                  3
                              )}) ${receiver.phone.substring(
                                  3,
                                  6
                              )}-${receiver.phone.substring(6, 10)}`
                            : receiver.phone}
                    </Span>
                </Box>
            ) : null}
            {receiver.type ? (
                <Box mt={4}>
                    <Span fontWeight="bold">Receiver Type: </Span>
                    <Span>{receiver.type}</Span>
                </Box>
            ) : null}
            {receiver.address &&
            receiver.city &&
            receiver.state &&
            receiver.zip ? (
                <Box mt={4}>
                    <Span fontWeight="bold">Receiver Address: </Span>
                    <Text mt={1}>{receiver.address}</Text>
                    <Text>{`${receiver.city}, ${receiver.state} ${
                        receiver.zip
                    }`}</Text>
                </Box>
            ) : null}
        </div>
    );
};

export default Receiver;
