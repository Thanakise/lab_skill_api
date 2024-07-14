import { Client } from "pg";

const client = new Client({
  user: "postgres",
  password: "postgres",
  host: "localhost",
  port: 5432,
  database: "app",
});

export const insertData = async () => {
  try {
    await client.connect();
    await client.query(`
              INSERT INTO
                    skill (key, name, description, tags)
                VALUES
                    (
                        'go',
                        'go',
                        'Go is an open source programming...',
                        '{go,golang}'
                    ) 
          `);
  } catch (e) {
    console.log(e);
  }
};
export const deleteData = async () => {
  try {
    await client.query(`
              DELETE FROM skill WHERE key = 'go'
          `);
    await client.end();
  } catch (e) {
    console.log(e);
  }
};
