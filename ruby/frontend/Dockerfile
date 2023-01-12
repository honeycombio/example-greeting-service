FROM ruby:2.7
RUN gem install bundler
WORKDIR /myapp
COPY Gemfile* /myapp/
RUN bundle install
COPY frontend.ru /myapp

EXPOSE 7007
CMD [ "bundle", "exec", "rackup", "frontend.ru", "--server", "puma", "--host", "0.0.0.0"]
