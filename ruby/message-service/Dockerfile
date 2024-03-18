FROM ruby:3.0
RUN gem install bundler
WORKDIR /myapp
COPY Gemfile /myapp/Gemfile
RUN bundle install
COPY message.ru /myapp

EXPOSE 9000
CMD [ "bundle", "exec", "rackup", "message.ru", "--server", "puma", "--host", "0.0.0.0"]
