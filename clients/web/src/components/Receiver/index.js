import React from "react";

import Span from "../Span";
import Text from "../Text";

const Receiver = ({ receiver }) => {
    return (
        <div>
            {receiver.name ? (
                <div>
                    <Span>{receiver.name}</Span>
                </div>
            ) : null}
            {receiver.phone ? (
                <div>
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
                </div>
            ) : null}
            {receiver.type ? (
                <div>
                    <Span>{receiver.type}</Span>
                </div>
            ) : null}
            {receiver.address &&
            receiver.city &&
            receiver.state &&
            receiver.zip ? (
                <div>
                    <Text mt={1}>{receiver.address}</Text>
                    <Text>{`${receiver.city}, ${receiver.state} ${
                        receiver.zip
                    }`}</Text>
                </div>
            ) : null}
        </div>
    );
};

export default Receiver;
