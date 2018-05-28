import React from "react";

import Box from "../Box";
import Span from "../Span";
import Text from "../Text";

const RadioReport = ({ radioReport }) => {
    return (
        <div>
            <Box mt={4}>
                <Span fontWeight="bold">Short report</Span>
                <Text mt={1}>{radioReport.shortReport}</Text>
            </Box>
            {radioReport.gender ? (
                <Box mt={4}>
                    <Span fontWeight="bold">Sex: </Span>
                    <Span>{radioReport.gender}</Span>
                </Box>
            ) : null}
            {radioReport.age ? (
                <Box mt={4}>
                    <Span fontWeight="bold">Age: </Span>
                    <Span>{radioReport.age}</Span>
                </Box>
            ) : null}
            {radioReport.weight ? (
                <Box mt={4}>
                    <Span fontWeight="bold">Weight: </Span>
                    <Span>{`${radioReport.weight} kg`}</Span>
                </Box>
            ) : null}
            <Box mt={4}>
                <Span fontWeight="bold">Intubated: </Span>
                <Span>{radioReport.intubated ? "Y" : "N"}</Span>
            </Box>
            <Box mt={4}>
                <Span fontWeight="bold">Cardiac: </Span>
                <Span>{radioReport.cardiac ? " Y" : " N"}</Span>
                <Span fontSize={20} px={2}>
                    |
                </Span>
                <Span fontWeight="bold">GIBleed: </Span>
                <Span>{radioReport.GIBleed ? " Y" : " N"}</Span>
                <Span fontSize={20} px={2}>
                    |
                </Span>
                <Span fontWeight="bold">OB: </Span>
                <Span>{radioReport.OB ? " Y" : " N"}</Span>
            </Box>
            {radioReport.drips ? (
                <Box mt={4}>
                    <Span fontWeight="bold">Drips: </Span>
                    <Span>{radioReport.drips}</Span>
                </Box>
            ) : null}
        </div>
    );
};

export default RadioReport;
