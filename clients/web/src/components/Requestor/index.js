import React from "react";

import Span from "../Span";
import Text from "../Text";

const Requestor = ({ requestor }) => {
    return (
        <div>
            {requestor.name ? (
                <div>
                    <Span>{requestor.name}</Span>
                </div>
            ) : null}
            {requestor.phone ? (
                <div>
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
                </div>
            ) : null}
            {requestor.type ? (
                <div>
                    <Span>{requestor.type}</Span>
                </div>
            ) : null}
            {requestor.address &&
            requestor.city &&
            requestor.state &&
            requestor.zip ? (
                <div>
                    <Text mt={1}>{requestor.address}</Text>
                    <Text>{`${requestor.city}, ${requestor.state} ${
                        requestor.zip
                    }`}</Text>
                </div>
            ) : null}
        </div>
    );
};

export default Requestor;
