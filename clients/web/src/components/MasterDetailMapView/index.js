import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import ScrollView from "../ScrollView";
import MasterDetailView from "../MasterDetailView";
import MasterView from "../MasterView";
import DetailView from "../DetailView";
import ReactMapboxGl from "react-mapbox-gl";

const Map = ReactMapboxGl({
    accessToken: process.env.REACT_APP_MAPBOX
});

const MasterDetailMapView = ({
    renderMasterView,
    renderDetailView,
    renderMapView,
    mapCenter,
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

                    <Flex flex={1}>
                        <Map
                            style="mapbox://styles/mapbox/streets-v9"
                            containerStyle={{
                                width: "100%",
                                height: "100%"
                            }}
                            center={mapCenter()}
                        >
                            {renderMapView()}
                        </Map>
                    </Flex>
                </Flex>
            </DetailView>
        </MasterDetailView>
    );
};

export default MasterDetailMapView;
