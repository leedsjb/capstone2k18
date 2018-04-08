import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchAircaftDetail } from "../../actions/aircraftDetail/actions";

class AircraftDetailProvider extends Component {
    componentDidMount() {
        this.props.fetchAircaftDetail(this.props.id);
    }

    render() {
        return this.props.render({ aircraftDetail: this.props.aircraftDetail });
    }
}

const mapStateToProps = ({ aircraftDetail }) => {
    return {
        aircraftDetail
    };
};

const mapDispatchToProps = {
    fetchAircaftDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(
    AircraftDetailProvider
);
