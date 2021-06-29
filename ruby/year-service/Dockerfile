FROM ruby:3.0
RUN gem install bundler
WORKDIR /app
COPY Gemfile /app/Gemfile
COPY Gemfile.lock /app/Gemfile.lock
RUN bundle install
COPY config.ru /app

EXPOSE 6001
CMD ["rackup", "--host", "0.0.0.0", "--port", "6001"]
