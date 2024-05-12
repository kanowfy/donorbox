import { Dropdown, Button } from "flowbite-react";
import ProjectCard from "./ProjectCard";

const SearchResults = () => {
  return (
    <div>
      <div className="flex gap-1">
        <Dropdown
          label="Location"
          dismissOnClick={false}
          color="light"
          outline
          size="sm"
          pill
        >
          <Dropdown.Header>
            <span>Choose location worldwide</span>
          </Dropdown.Header>
          <Dropdown.Item>
            <form>
              <input></input>
            </form>
          </Dropdown.Item>
        </Dropdown>
        <Dropdown
          label="Category"
          dismissOnClick={false}
          color="light"
          outline
          size="sm"
          pill
        >
          <Dropdown.Header>
            <span>Choose a category</span>
          </Dropdown.Header>
          <Dropdown.Item>
            <div className="grid grid-cols-4 max-w-sm gap-2">
              <Button outline size="xs" color="success" pill>
                Medical
              </Button>
              <Button outline size="xs" color="success" pill>
                Emergency
              </Button>
              <Button outline size="xs" color="success" pill>
                Events
              </Button>
              <Button outline size="xs" color="success" pill>
                Non profit
              </Button>
            </div>
          </Dropdown.Item>
        </Dropdown>
        <Button outline size="sm" color="light" pill>
          Close to goal
        </Button>
        <Button outline size="sm" color="light" pill>
          Recently launched
        </Button>
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
  );
};

export default SearchResults;
