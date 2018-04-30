import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import DetailView from "../DetailView";
import MapView from "../MapView";
import MasterView from "../MasterView";
import MasterDetailView from "../MasterDetailView";
import ScrollView from "../ScrollView";

const MasterDetailMapView = ({
    renderMasterView,
    renderDetailView,
    showDetail
}) => {
    return (
        <MasterDetailView>
            <MasterView>{renderMasterView()}</MasterView>
            <DetailView>
                <Flex
                    style={{ height: "100%", width: "100%" }}
                    flexDirection={[null, null, "column", "row"]}
                >
                    {showDetail ? (
                        <ScrollView
                            maxWidth={[null, null, null, 320]}
                            height={[null, null, "50%", "100%"]}
                        >
                            <Box p={3}>{renderDetailView()}</Box>
                        </ScrollView>
                    ) : null}

                    <MapView id={showDetail} />
                </Flex>
            </DetailView>
        </MasterDetailView>
    );
};

export default MasterDetailMapView;
