import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchAircraft } from "../../actions/aircraft/actions";

class AircraftProvider extends Component {
    componentDidMount() {
        this.props.fetchAircraft(this.props.status);
    }
    // TODO: Update with new syntax
    componentWillReceiveProps(nextProps) {
        console.log("STATUS", this.props.status);
        if (nextProps !== this.props) {
            this.props.fetchAircraft(this.props.status);
        }
    }

    render() {
        console.log("STATUS", this.props.status);
        return this.props.render({ aircraft: this.props.aircraft });
    }
}

const mapStateToProps = ({ aircraft }) => {
    return {
        aircraft
    };
};

const mapDispatchToProps = {
    fetchAircraft
};

export default connect(mapStateToProps, mapDispatchToProps)(AircraftProvider);
