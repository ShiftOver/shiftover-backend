require('dotenv').config();
require('dotenv').config({ path: '.env.secrets' });

const config = {
  mongodb: {
    url: `mongodb+srv://${process.env.MONGO_USERNAME}:${process.env.MONGO_PASSWORD}@${process.env.MONGO_HOST}`,
    databaseName: process.env.MONGO_DBNAME,
  },
  migrationsDir: "migrations",
  changelogCollectionName: "changelog",
  migrationFileExtension: ".js",
  useFileHash: false,
  moduleSystem: 'commonjs',
};

module.exports = config;
