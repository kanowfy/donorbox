import { DonutChart, Legend } from "@tremor/react";
import PropTypes from "prop-types";
import { useEffect, useState } from "react";

const DashboardPieChart = ({ stats }) => {
  const [statusData, setStatusData] = useState();

  useEffect(() => {
    const constructStatusData = (stats) => {
      setStatusData([
        {
          name: "Ongoing",
          quantity: stats?.ongoing,
        },
        {
          name: "Completed Payout",
          quantity: stats?.completed_payout,
        },
        {
          name: "Completed Refund",
          quantity: stats?.completed_refund,
        },
        {
          name: "Pending Action or No Fund",
          quantity: stats?.ended,
        },
      ]);
    };

    if (stats) {
      constructStatusData(stats);
    }
  }, [stats]);

  return (
    <div className="bg-white rounded-xl py-10 pl-10 w-fit space-y-5">
      <div className="text-base font-semibold">Fundraiser Distribution</div>
      <div className="flex items-center justify-center space-x-6">
        <DonutChart
          data={statusData}
          category="quantity"
          index="name"
          colors={["blue", "indigo", "fuchsia", "teal"]}
          className="w-60 h-60"
          variant="pie"
        />
        <Legend
          categories={[
            "Ongoing",
            "Completed Payout",
            "Completed Refund",
            "Pending Action or No Fund",
          ]}
          colors={["blue", "indigo", "fuchsia", "teal"]}
          className="max-w-xs"
        />
      </div>
    </div>
  );
};

DashboardPieChart.propTypes = {
  stats: PropTypes.object,
};

export default DashboardPieChart;
