import { test, expect } from "@playwright/test";
import { deleteData, insertData } from "../database/initdata";

test.beforeAll(() => insertData());
test.afterAll(() => deleteData());
test("get single skill", async ({ request }) => {
  const skill = await request.get(`/api/v1/skills/go`);

  expect(skill.ok()).toBeTruthy();
  expect(await skill.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "go",
        name: "go",
        description: expect.any(String),
        logo: expect.any(String),
        tags: expect.any(Array<String>),
      },
    })
  );
  const error = await request.get(`/api/v1/skills/${Math.random()}`);
  expect(!error.ok()).toBeTruthy();
  expect(await error.json()).toEqual(
    expect.objectContaining({
      status: "error",
      message: "Skill not found",
    })
  );
});
test("get all skills", async ({ request }) => {
  const skill = await request.get(`/api/v1/skills`);

  expect(skill.ok()).toBeTruthy();
  expect(await skill.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: expect.arrayContaining([
        {
          key: "go",
          name: "go",
          description: expect.any(String),
          logo: expect.any(String),
          tags: expect.any(Array<String>),
        },
      ]),
    })
  );
});
test("post skill", async ({ request }) => {
  let expectData = {
    key: "test",
    name: "test",
    logo: "test",
    description: "test",
    tags: ["test", "test"],
  };
  const skill = await request.post(`/api/v1/skills`, {
    data: expectData,
  });

  expect(skill.ok()).toBeTruthy();
  expect(await skill.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: expectData,
    })
  );

  const newSkill = await request.post(`/api/v1/skills`, {
    data: expectData,
  });
  expect(!newSkill.ok()).toBeTruthy();
  expect(await newSkill.json()).toEqual(
    expect.objectContaining({
      status: "error",
      message: "Skill already exists",
    })
  );

  await request.delete(`/api/v1/skills/test`);
});
test("put skill", async ({ request }) => {
  let expectData = {
    name: "Python 3",
    description:
      "Python 3 is the latest version of Python programming language.",
    logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
    tags: ["data"],
  };
  const skill = await request.put(`/api/v1/skills/go`, {
    data: expectData,
  });

  expect(skill.ok()).toBeTruthy();
  expect(await skill.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "go",
        ...expectData,
      },
    })
  );
  const error = await request.put(`/api/v1/skills/${Math.random()}`, {
    data: expectData,
  });
  expect(!error.ok()).toBeTruthy();
  expect(await error.json()).toEqual(
    expect.objectContaining({
      status: "error",
      message: "not be able to update skill",
    })
  );
});
test("patch skill name", async ({ request }) => {
  let expectData = {
    name: "Python 3",
  };
  const skill = await request.patch(`/api/v1/skills/go/actions/name`, {
    data: expectData,
  });

  expect(skill.ok()).toBeTruthy();
  expect(await skill.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "go",
        name: expectData.name,
        description: expect.any(String),
        logo: expect.any(String),
        tags: expect.any(Array<String>),
      },
    })
  );

  const error = await request.patch(
    `/api/v1/skills/${Math.random()}/actions/name`,
    {
      data: expectData,
    }
  );
  expect(!error.ok()).toBeTruthy();
  expect(await error.json()).toEqual(
    expect.objectContaining({
      status: "error",
      message: "not be able to update skill name",
    })
  );
});
test("patch skill description", async ({ request }) => {
  let expectData = {
    description:
      "Python 3 is the latest version of Python programming language.",
  };

  const skill = await request.patch(`/api/v1/skills/go/actions/description`, {
    data: expectData,
  });

  expect(skill.ok()).toBeTruthy();
  expect(await skill.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "go",
        name: expect.any(String),
        description: expectData.description,
        logo: expect.any(String),
        tags: expect.any(Array<String>),
      },
    })
  );
  const error = await request.patch(
    `/api/v1/skills/${Math.random()}/actions/description`,
    {
      data: expectData,
    }
  );
  expect(!error.ok()).toBeTruthy();
  expect(await error.json()).toEqual(
    expect.objectContaining({
      status: "error",
      message: "not be able to update skill description",
    })
  );
});
test("patch skill logo", async ({ request }) => {
  let expectData = {
    logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
  };

  const skill = await request.patch(`/api/v1/skills/go/actions/logo`, {
    data: expectData,
  });

  expect(skill.ok()).toBeTruthy();
  expect(await skill.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "go",
        name: expect.any(String),
        description: expect.any(String),
        logo: expectData.logo,
        tags: expect.any(Array<String>),
      },
    })
  );
  const error = await request.patch(
    `/api/v1/skills/${Math.random()}/actions/logo`,
    {
      data: expectData,
    }
  );
  expect(!error.ok()).toBeTruthy();
  expect(await error.json()).toEqual(
    expect.objectContaining({
      status: "error",
      message: "not be able to update skill logo",
    })
  );
});
test("patch skill tags", async ({ request }) => {
  let expectData = {
    tags: ["programming language", "data"],
  };

  const skill = await request.patch(`/api/v1/skills/go/actions/tags`, {
    data: expectData,
  });

  expect(skill.ok()).toBeTruthy();
  expect(await skill.json()).toEqual(
    expect.objectContaining({
      status: "success",
      data: {
        key: "go",
        name: expect.any(String),
        description: expect.any(String),
        logo: expect.any(String),
        tags: expectData.tags,
      },
    })
  );
  const error = await request.patch(
    `/api/v1/skills/${Math.random()}/actions/tags`,
    {
      data: expectData,
    }
  );
  expect(!error.ok()).toBeTruthy();
  expect(await error.json()).toEqual(
    expect.objectContaining({
      status: "error",
      message: "not be able to update skill tags",
    })
  );
});
test("delete skill", async ({ request }) => {
  const skill = await request.delete(`/api/v1/skills/go`);

  expect(skill.ok()).toBeTruthy();
  expect(await skill.json()).toEqual(
    expect.objectContaining({
      status: "success",
      message: "Skill deleted",
    })
  );
  const error = await request.delete(`/api/v1/skills/${Math.random()}`);
  expect(!error.ok()).toBeTruthy();
  expect(await error.json()).toEqual(
    expect.objectContaining({
      status: "error",
      message: "not be able to delete skill",
    })
  );
});
