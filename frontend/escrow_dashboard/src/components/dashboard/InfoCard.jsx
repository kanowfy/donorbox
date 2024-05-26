import { Card } from "@tremor/react";
import PropTypes from "prop-types";

const InfoCard = ({ title, content }) => {
  return (
    <Card className="mx-auto space-y-4">
      <div className="text-lg text-tremor-content dark:text-dark-tremor-content">
        {title}
      </div>
      <p className="text-tremor-metric font-semibold text-tremor-content-strong dark:text-dark-tremor-content-strong tracking-tight">
        {content}
      </p>
    </Card>
  );
};

InfoCard.propTypes = {
  title: PropTypes.string,
  content: PropTypes.string,
};

export default InfoCard;
