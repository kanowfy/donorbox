import { LineChart } from "@tremor/react";
import PropTypes from "prop-types";
import { useEffect, useState } from "react";

const DashboardLineChart = ({ transactions }) => {
  const [chartData, setChartData] = useState();

  useEffect(() => {
    const constructChartData = (transactions) => {
      setChartData(
        transactions?.map((t) => {
          return {
            Week: t.week,
            Backings: t.backings,
            Payouts: t.payouts,
            Refunds: t.refunds,
          };
        })
      );
    };

    if (transactions) {
      constructChartData(transactions);
    }
  }, [transactions]);

  return (
    <div className="bg-white rounded-xl py-10 px-10 w-full space-y-5">
      <div className="text-base font-semibold">Total weekly transactions</div>
      <div className="flex items-center justify-center">
        <LineChart
          className="mt-4 h-72"
          data={chartData}
          index="Week"
          yAxisWidth={65}
          categories={["Backings", "Payouts", "Refunds"]}
          colors={["indigo", "cyan", "fuchsia"]}
          xAxisLabel="Week"
          yAxisLabel="Number of Transactions"
        />
      </div>
    </div>
  );
};

DashboardLineChart.propTypes = {
  transactions: PropTypes.array,
};

export default DashboardLineChart;
