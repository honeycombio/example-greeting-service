FROM ruby:3.0
RUN gem install bundler
WORKDIR /app
COPY Gemfile /app/Gemfile
COPY Gemfile.lock /app/Gemfile.lock
RUN bundle install
COPY config.ru /app
COPY o11y_wrapper.rb /app

EXPOSE 6001
CMD ["bundle", "exec", "puma", "--port", "6001"]
