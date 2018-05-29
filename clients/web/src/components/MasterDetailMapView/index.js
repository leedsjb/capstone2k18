import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import DetailView from "../DetailView";
import MapView from "../MapView";
import MasterView from "../MasterView";
import MasterDetailView from "../MasterDetailView";

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
                        <Box
                            maxWidth={[null, null, null, 400]}
                            width={1}
                            height={[null, null, "50%", "100%"]}
                            style={{
                                display: "flex",
                                flexDirection: "column"
                            }}
                        >
                            {renderDetailView()}
                        </Box>
                    ) : null}

                    <MapView id={showDetail} />
                </Flex>
            </DetailView>
        </MasterDetailView>
    );
};

export default MasterDetailMapView;
