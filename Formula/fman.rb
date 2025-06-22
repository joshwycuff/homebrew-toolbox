class Fman < Formula
  desc "Perform fuzzy search on man pages."
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/fman-237b2269bdd98d036a5120278c2b6897ad47df87/fman",
  version "237b2269bdd98d036a5120278c2b6897ad47df87"

  depends_on "apropos"
  depends_on "fzf"

  def install
    bin.install "tools/fman/fman.sh" => "fman"
  end
end
