import React from "react";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";

import FlexFullHeight from "../../components/FlexFullHeight";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import ScrollView from "../../components/ScrollView";
import MasterDetailView from "../../components/MasterDetailView";
import MasterView from "../../components/MasterView";
import DetailView from "../../components/DetailView";
import Avatar from "../../components/Avatar";
import Heading from "../../components/Heading";
import ButtonIcon from "../../components/ButtonIcon";
import NavBar from "../../components/NavBar";

const PeoplePage = () => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet>
                <title>Missions</title>
            </Helmet>

            <TitleBar title="People" />
            <NavBar />

            <MasterDetailView>
                <MasterView>
                    Test 123<br />
                    Test 123
                </MasterView>
                <DetailView>
                    <Flex flexDirection="column">
                        <Avatar
                            src="https://pbs.twimg.com/profile_images/867506462142156800/rlT9Ppkp_400x400.jpg"
                            size={96}
                        />
                        <Heading children="Vincent van der Meulen" is="h2" />
                        <Heading children="Maker" is="h3" />
                        <Flex>
                            <ButtonIcon>Text</ButtonIcon>
                            <ButtonIcon>Call</ButtonIcon>
                            <ButtonIcon>Mail</ButtonIcon>
                        </Flex>
                    </Flex>
                </DetailView>
            </MasterDetailView>
            <TabBar />
        </FlexFullHeight>
    );
};

export default PeoplePage;
