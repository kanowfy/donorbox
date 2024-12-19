import { Link } from "react-router-dom";
import ProjectCard from "../components/ProjectCard";
import { useEffect, useState } from "react";
import projectService from "../services/project";
import { FaSearch } from "react-icons/fa";
import utils from "../utils/utils";

const Home = () => {
  const [ongoingProjects, setOngoingProjects] = useState([]);
  const [successfulProjects, setSuccessfulProjects] = useState([]);
  const [displayOngoing, setDisplayOngoing] = useState(true);

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const resp = await projectService.getAll();
        console.log(resp.projects);
        setOngoingProjects(resp.projects.filter((p) => p.status === "ongoing"));
        setSuccessfulProjects(
          resp.projects.filter((p) => p.status === "finished")
        );
        /*
        setSuccessfulProjects(
          resp.projects.filter((p) => p.status === "finished")
        );
        */
      } catch (err) {
        console.error(err);
      }
    };

    fetchProjects();
  }, []);

  return (
    <>
      <div>
        <section className="relative h-[48rem]">
          <div
            className="absolute inset-0 bg-top bg-cover bg-no-repeat filter brightness-50"
            style={{ backgroundImage: "url('/handshake.jpeg')" }}
          ></div>

          <div className="relative z-10 flex items-center flex-col top-36">
            <div className="pb-7 font-semibold text-gray-100 text-7xl">
              Help those in need today
            </div>
            <div className="text-gray-100 font-medium text-2xl">
              Your home for communities, charities and people you care about
            </div>

            {/*<div>
            <Link to="/start-fundraiser">
              <button className="text-lg mt-10 px-8 py-4 text-white font-semibold rounded-xl shadow-lg bg-gradient-to-t from-cyan-500 to-blue-500 hover:bg-gradient-to-b">
                Start a Fundraiser
              </button>
            </Link>
          </div>*/}
          </div>
        </section>

        <section>
          <div className="min-h-screen px-10 pt-6 bg-gray-50">
            <div className="flex justify-center mx-10 gap-5">
              <button className={`text-xl border border-gray-300 ${displayOngoing && "bg-gray-300"} text-gray-700 px-5 py-3 rounded-full shadow-sm hover:shadow-lg`}
              onClick={() => setDisplayOngoing(true)}>
                Ongoing Fundraisers
              </button>
              <button className={`text-xl border border-gray-300 ${!displayOngoing && "bg-gray-300"} px-5 py-3 rounded-full text-gray-600 shadow-sm hover:shadow-lg`}
              onClick={() => setDisplayOngoing(false)}>
                Finished Fundraisers
              </button>
            </div>
            <div className="flex justify-center mb-20">
              {displayOngoing ? (
                <div className="grid grid-cols-1 gap-10 md:grid-cols-2 xl:grid-cols-3 mx-10 mt-10">
                  {ongoingProjects &&
                    ongoingProjects
                      .slice(0, 9)
                      .map((p) => (
                        <ProjectCard
                          id={p.id}
                          title={p.title}
                          cover={p.cover_picture}
                          totalFund={p.total_fund}
                          fundGoal={p.fund_goal}
                          numBackings={p.backing_count}
                          category={utils.getCategoryNameByID(p.category_id)}
                          key={p.id}
                        />
                      ))}
                </div>
              ) : (
                <div className="grid grid-cols-1 gap-10 md:grid-cols-2 xl:grid-cols-3 mx-10 mt-10">
                  {successfulProjects &&
                    successfulProjects
                      .slice(0, 9)
                      .map((p) => (
                        <ProjectCard
                          id={p.id}
                          title={p.title}
                          cover={p.cover_picture}
                          totalFund={p.total_fund}
                          fundGoal={p.fund_goal}
                          numBackings={p.backing_count}
                          category={utils.getCategoryNameByID(p.category_id)}
                          key={p.id}
                        />
                      ))}
                </div>
              )}
            </div>
          </div>
        </section>
        {/*successfulProjects?.length > 0 && (
            <section>
              <div className="px-10 pt-6 bg-gray-50">
                <div className="flex justify-between mx-48">
                  <div className="font-medium text-2xl tracking-tight">
                    Successful fundraisers
                  </div>
                </div>
                <div className="flex justify-center">
                  <div className="grid grid-cols-1 gap-7 md:grid-cols-2 xl:grid-cols-4 mx-16 mb-10">
                    {successfulProjects.slice(0, 8).map((p) => (
                      <ProjectCard
                        id={p.id}
                        title={p.title}
                        cover={p.cover_picture}
                        totalFund={p.total_fund}
                        fundGoal={p.fund_goal}
                        numBackings={p.backing_count}
                        key={p.id}
                      />
                    ))}
                  </div>
                </div>
              </div>
            </section>
          )*/}
      </div>
    </>
  );
};

export default Home;
