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
    renderMapView
}) => {
    return (
        <MasterDetailView>
            <MasterView>{renderMasterView()}</MasterView>
            <DetailView>
                <Flex
                    style={{ height: "100%", width: "100%" }}
                    flexDirection={[null, null, "column", "row"]}
                >
                    <ScrollView
                        maxWidth={[null, null, null, 320]}
                        height={[null, null, "50%", "100%"]}
                    >
                        <Box p={3}>{renderDetailView()}</Box>
                    </ScrollView>

                    <Flex flex={1}>
                        <Map
                            style="mapbox://styles/mapbox/streets-v9"
                            containerStyle={{
                                width: "100%",
                                height: "100%"
                            }}
                            center={[6.5665, 53.2194]}
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
