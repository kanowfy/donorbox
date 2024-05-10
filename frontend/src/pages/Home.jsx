import { Button } from "flowbite-react";
import { Link } from "react-router-dom";
import ProjectCard from "../components/ProjectCard";

const Home = () => {
  return (
    <>
      <div>
        <section className="flex flex-col pt-20 items-center h-128 bg-cover bg-gradient-to-b from-white to-emerald-300">
          <div className="pb-7 font-semibold text-green-900 text-6xl">
            Help those in need today
          </div>
          <div className="text-green-700 pb-10 font-medium text-lg">
            Your home for communities, charities and people you care about
          </div>
          <div>
            <Button
              gradientDuoTone="greenToBlue"
              pill
              size="xl"
              className="mt-5"
            >
              Start a Fundraiser
            </Button>
          </div>
        </section>

        <section>
          <div className="min-h-screen px-10 pt-6">
            <div className="flex justify-between mx-48">
              <div className="font-medium text-2xl tracking-tight">
                Popular fundraisers right now
              </div>
              <div>
                <Link to="#">
                  <div className="underline font-semibold text-xl text-gray-800 hover:text-green-800">
                    Explore
                  </div>
                </Link>
              </div>
            </div>
            <div className="flex justify-center">
              <div className="grid grid-cols-1 gap-7 md:grid-cols-3 xl:grid-cols-4 mx-16 mt-10 mb-16">
                <div>
                  <ProjectCard
                    title="Celebration of my homie lorem ipsum skibidi toilet rizz edge goon"
                    cover="https://w.wallhaven.cc/full/l8/wallhaven-l8vp7y.jpg"
                    currentAmount={1000}
                    goalAmount={3000}
                    numBackings={20}
                  />
                </div>
                <div>
                  <ProjectCard
                    title="Celebration of my homie"
                    cover="https://w.wallhaven.cc/full/l8/wallhaven-l8vp7y.jpg"
                    currentAmount={1000}
                    goalAmount={3000}
                    numBackings={20}
                  />
                </div>
                <div>
                  <ProjectCard
                    title="Celebration of my homie"
                    cover="https://static.vecteezy.com/system/resources/thumbnails/025/284/015/small_2x/close-up-growing-beautiful-forest-in-glass-ball-and-flying-butterflies-in-nature-outdoors-spring-season-concept-generative-ai-photo.jpg"
                    currentAmount={2000}
                    goalAmount={3000}
                    numBackings={20}
                  />
                </div>
                <div>
                  <ProjectCard
                    title="Celebration of my homie"
                    cover="https://static.vecteezy.com/system/resources/thumbnails/025/284/015/small_2x/close-up-growing-beautiful-forest-in-glass-ball-and-flying-butterflies-in-nature-outdoors-spring-season-concept-generative-ai-photo.jpg"
                    currentAmount={2500}
                    goalAmount={3000}
                    numBackings={20}
                  />
                </div>
                <div>
                  <ProjectCard
                    title="Celebration of my homie"
                    cover="https://w.wallhaven.cc/full/l8/wallhaven-l8vp7y.jpg"
                    currentAmount={2780}
                    goalAmount={3000}
                    numBackings={20}
                  />
                </div>
                <div>
                  <ProjectCard
                    title="Celebration of my homie"
                    cover="https://w.wallhaven.cc/full/l8/wallhaven-l8vp7y.jpg"
                    currentAmount={1000}
                    goalAmount={3000}
                    numBackings={20}
                  />
                </div>
              </div>
            </div>
          </div>
        </section>
      </div>
    </>
  );
};

export default Home;
