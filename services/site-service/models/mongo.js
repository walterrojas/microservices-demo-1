const Promise = require('bluebird')
const mongoose = require('mongoose')

const Logger = require('../util/logger')

class Mongo {
  constructor (config, options) {
    options = options || {}
    this.mongoUri = config.mongoUri
    this.mongoUsername = config.mongoUsername
    this.mongoPassword = config.mongoPassword
    this.logger = options.logger || new Logger('MongoProvider')
    this.mongoose = options.mongoose || mongoose
    this.mongoose.Promise = global.Promise
  }

  connect () {
    return new Promise((resolve, reject) => {
      let connection = this.mongoose.connection

      let opts = {
        useNewUrlParser: true,
        autoReconnect: true
      }

      if (this.mongoUsername && this.mongoPassword) {
        opts.auth = {
          user: this.mongoUsername,
          password: this.mongoPassword
        }
      }

      this.mongoose.connect(this.mongoUri, opts).catch(err => {
        this.logger.error('Error connecting to Mongo.', err)
        reject(err)
      })

      const errorListener = err => {
        this.logger.error('Error connecting to Mongo.', err)
        reject(err)
      }

      connection.once('error', errorListener)

      connection.on('connected', () => {
        this.logger.info('Connected to Mongo.')
      })

      connection.on('disconnected', () => {
        this.logger.warn('Disocnnected from Mongo.')
      })

      connection.on('open', () => {
        this.logger.info('Mongo connection is open.')
        connection.removeListener('error', errorListener)
        resolve(connection)
      })

      connection.on('close', () => {
        this.logger.error('Mongo connection is closed.')
      })
    })
  }

  disconnect (callback) {
    const promise = this.mongoose.disconnect(callback)
    promise.catch(err => this.logger.error('Error disconnecting from Mongo.', err))
    return promise
  }
}

module.exports = Mongo
